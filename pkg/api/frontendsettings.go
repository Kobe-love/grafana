package api

import (
	"context"
	"strconv"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/plugins"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/tsdb/grafanads"
	"github.com/grafana/grafana/pkg/util"
)

func (hs *HTTPServer) getFSDataSources(c *models.ReqContext, enabledPlugins EnabledPlugins) (map[string]plugins.DataSourceDTO, error) {
	orgDataSources := make([]*models.DataSource, 0)

	if c.OrgId != 0 {
		query := models.GetDataSourcesQuery{OrgId: c.OrgId, DataSourceLimit: hs.Cfg.DataSourceLimit}
		err := bus.Dispatch(c.Req.Context(), &query)

		if err != nil {
			return nil, err
		}

		filtered, err := filterDatasourcesByQueryPermission(c.Req.Context(), c.SignedInUser, query.Result)
		if err != nil {
			return nil, err
		}

		orgDataSources = filtered
	}

	dataSources := make(map[string]plugins.DataSourceDTO)

	for _, ds := range orgDataSources {
		url := ds.Url

		if ds.Access == models.DS_ACCESS_PROXY {
			url = "/api/datasources/proxy/" + strconv.FormatInt(ds.Id, 10)
		}

		dsDTO := plugins.DataSourceDTO{
			ID:        ds.Id,
			UID:       ds.Uid,
			Type:      ds.Type,
			Name:      ds.Name,
			URL:       url,
			IsDefault: ds.IsDefault,
			Access:    string(ds.Access),
		}

		plugin, exists := enabledPlugins.Get(plugins.DataSource, ds.Type)
		if !exists {
			c.Logger.Error("Could not find plugin definition for data source", "datasource_type", ds.Type)
			continue
		}
		dsDTO.Preload = plugin.Preload
		dsDTO.Module = plugin.Module
		dsDTO.PluginMeta = &plugins.PluginMetaDTO{
			JSONData:  plugin.JSONData,
			Signature: plugin.Signature,
			Module:    plugin.Module,
			BaseURL:   plugin.BaseURL,
		}

		if ds.JsonData == nil {
			dsDTO.JSONData = make(map[string]interface{})
		} else {
			dsDTO.JSONData = ds.JsonData.MustMap()
		}

		if ds.Access == models.DS_ACCESS_DIRECT {
			if ds.BasicAuth {
				dsDTO.BasicAuth = util.GetBasicAuthHeader(
					ds.BasicAuthUser,
					hs.DataSourcesService.DecryptedBasicAuthPassword(ds),
				)
			}
			if ds.WithCredentials {
				dsDTO.WithCredentials = ds.WithCredentials
			}

			if ds.Type == models.DS_INFLUXDB_08 {
				dsDTO.Username = ds.User
				dsDTO.Password = hs.DataSourcesService.DecryptedPassword(ds)
				dsDTO.URL = url + "/db/" + ds.Database
			}

			if ds.Type == models.DS_INFLUXDB {
				dsDTO.Username = ds.User
				dsDTO.Password = hs.DataSourcesService.DecryptedPassword(ds)
				dsDTO.URL = url
			}
		}

		if (ds.Type == models.DS_INFLUXDB) || (ds.Type == models.DS_ES) {
			dsDTO.Database = ds.Database
		}

		if ds.Type == models.DS_PROMETHEUS {
			// add unproxied server URL for link to Prometheus web UI
			ds.JsonData.Set("directUrl", ds.Url)
		}

		dataSources[ds.Name] = dsDTO
	}

	// add data sources that are built in (meaning they are not added via data sources page, nor have any entry in
	// the datasource table)
	for _, ds := range hs.pluginStore.Plugins(c.Req.Context(), plugins.DataSource) {
		if ds.BuiltIn {
			dto := plugins.DataSourceDTO{
				Type:     string(ds.Type),
				Name:     ds.Name,
				JSONData: make(map[string]interface{}),
				PluginMeta: &plugins.PluginMetaDTO{
					JSONData:  ds.JSONData,
					Signature: ds.Signature,
					Module:    ds.Module,
					BaseURL:   ds.BaseURL,
				},
			}
			if ds.Name == grafanads.DatasourceName {
				dto.ID = grafanads.DatasourceID
				dto.UID = grafanads.DatasourceUID
			}
			dataSources[ds.Name] = dto
		}
	}

	return dataSources, nil
}

// getFrontendSettingsMap returns a json object with all the settings needed for front end initialisation.
func (hs *HTTPServer) getFrontendSettingsMap(c *models.ReqContext) (map[string]interface{}, error) {
	enabledPlugins, err := hs.enabledPlugins(c.Req.Context(), c.OrgId)
	if err != nil {
		return nil, err
	}

	pluginsToPreload := make([]*plugins.PreloadPlugin, 0)
	for _, app := range enabledPlugins[plugins.App] {
		if app.Preload {
			pluginsToPreload = append(pluginsToPreload, &plugins.PreloadPlugin{
				Path:    app.Module,
				Version: app.Info.Version,
			})
		}
	}

	dataSources, err := hs.getFSDataSources(c, enabledPlugins)
	if err != nil {
		return nil, err
	}

	defaultDS := "-- Grafana --"
	for n, ds := range dataSources {
		if ds.IsDefault {
			defaultDS = n
		}
	}

	panels := make(map[string]plugins.PanelDTO)
	for _, panel := range enabledPlugins[plugins.Panel] {
		if panel.State == plugins.AlphaRelease && !hs.Cfg.PluginsEnableAlpha {
			continue
		}

		panels[panel.ID] = plugins.PanelDTO{
			ID:            panel.ID,
			Name:          panel.Name,
			Info:          panel.Info,
			Module:        panel.Module,
			BaseURL:       panel.BaseURL,
			SkipDataQuery: panel.SkipDataQuery,
			HideFromList:  panel.HideFromList,
			ReleaseState:  string(panel.State),
			Signature:     string(panel.Signature),
			Sort:          getPanelSort(panel.ID),
		}
	}

	hideVersion := hs.Cfg.AnonymousHideVersion && !c.IsSignedIn
	version := setting.BuildVersion
	commit := setting.BuildCommit
	buildstamp := setting.BuildStamp

	if hideVersion {
		version = ""
		commit = ""
		buildstamp = 0
	}

	hasAccess := accesscontrol.HasAccess(hs.AccessControl, c)

	jsonObj := map[string]interface{}{
		"defaultDatasource":                   defaultDS,
		"datasources":                         dataSources,
		"minRefreshInterval":                  setting.MinRefreshInterval,
		"panels":                              panels,
		"appUrl":                              hs.Cfg.AppURL,
		"appSubUrl":                           hs.Cfg.AppSubURL,
		"allowOrgCreate":                      (setting.AllowUserOrgCreate && c.IsSignedIn) || c.IsGrafanaAdmin,
		"authProxyEnabled":                    setting.AuthProxyEnabled,
		"ldapEnabled":                         hs.Cfg.LDAPEnabled,
		"alertingEnabled":                     setting.AlertingEnabled,
		"alertingErrorOrTimeout":              setting.AlertingErrorOrTimeout,
		"alertingNoDataOrNullValues":          setting.AlertingNoDataOrNullValues,
		"alertingMinInterval":                 setting.AlertingMinInterval,
		"liveEnabled":                         hs.Cfg.LiveMaxConnections != 0,
		"autoAssignOrg":                       setting.AutoAssignOrg,
		"verifyEmailEnabled":                  setting.VerifyEmailEnabled,
		"sigV4AuthEnabled":                    setting.SigV4AuthEnabled,
		"exploreEnabled":                      setting.ExploreEnabled,
		"googleAnalyticsId":                   setting.GoogleAnalyticsId,
		"rudderstackWriteKey":                 setting.RudderstackWriteKey,
		"rudderstackDataPlaneUrl":             setting.RudderstackDataPlaneUrl,
		"rudderstackSdkUrl":                   setting.RudderstackSdkUrl,
		"rudderstackConfigUrl":                setting.RudderstackConfigUrl,
		"applicationInsightsConnectionString": hs.Cfg.ApplicationInsightsConnectionString,
		"applicationInsightsEndpointUrl":      hs.Cfg.ApplicationInsightsEndpointUrl,
		"disableLoginForm":                    setting.DisableLoginForm,
		"disableUserSignUp":                   !setting.AllowUserSignUp,
		"loginHint":                           setting.LoginHint,
		"passwordHint":                        setting.PasswordHint,
		"externalUserMngInfo":                 setting.ExternalUserMngInfo,
		"externalUserMngLinkUrl":              setting.ExternalUserMngLinkUrl,
		"externalUserMngLinkName":             setting.ExternalUserMngLinkName,
		"viewersCanEdit":                      setting.ViewersCanEdit,
		"editorsCanAdmin":                     hs.Cfg.EditorsCanAdmin,
		"disableSanitizeHtml":                 hs.Cfg.DisableSanitizeHtml,
		"pluginsToPreload":                    pluginsToPreload,
		"buildInfo": map[string]interface{}{
			"hideVersion":   hideVersion,
			"version":       version,
			"commit":        commit,
			"buildstamp":    buildstamp,
			"edition":       hs.License.Edition(),
			"latestVersion": hs.updateChecker.LatestGrafanaVersion(),
			"hasUpdate":     hs.updateChecker.GrafanaUpdateAvailable(),
			"env":           setting.Env,
		},
		"licenseInfo": map[string]interface{}{
			"expiry":          hs.License.Expiry(),
			"stateInfo":       hs.License.StateInfo(),
			"licenseUrl":      hs.License.LicenseURL(hasAccess(accesscontrol.ReqGrafanaAdmin, accesscontrol.LicensingPageReaderAccess)),
			"edition":         hs.License.Edition(),
			"enabledFeatures": hs.License.EnabledFeatures(),
		},
		"featureToggles":                   hs.Cfg.FeatureToggles,
		"rendererAvailable":                hs.RenderService.IsAvailable(),
		"rendererVersion":                  hs.RenderService.Version(),
		"http2Enabled":                     hs.Cfg.Protocol == setting.HTTP2Scheme,
		"sentry":                           hs.Cfg.Sentry,
		"pluginCatalogURL":                 hs.Cfg.PluginCatalogURL,
		"pluginAdminEnabled":               hs.Cfg.PluginAdminEnabled,
		"pluginAdminExternalManageEnabled": hs.Cfg.PluginAdminEnabled && hs.Cfg.PluginAdminExternalManageEnabled,
		"pluginCatalogHiddenPlugins":       hs.Cfg.PluginCatalogHiddenPlugins,
		"expressionsEnabled":               hs.Cfg.ExpressionsEnabled,
		"awsAllowedAuthProviders":          hs.Cfg.AWSAllowedAuthProviders,
		"awsAssumeRoleEnabled":             hs.Cfg.AWSAssumeRoleEnabled,
		"azure": map[string]interface{}{
			"cloud":                  hs.Cfg.Azure.Cloud,
			"managedIdentityEnabled": hs.Cfg.Azure.ManagedIdentityEnabled,
		},
		"caching": map[string]bool{
			"enabled": hs.Cfg.SectionWithEnvOverrides("caching").Key("enabled").MustBool(true),
		},
		"recordedQueries": map[string]bool{
			"enabled": hs.Cfg.SectionWithEnvOverrides("recorded_queries").Key("enabled").MustBool(false),
		},
		"unifiedAlertingEnabled": hs.Cfg.UnifiedAlerting.Enabled,
		"featureHighlights": map[string]bool{
			"enabled": hs.SettingsProvider.Section("feature_highlights").KeyValue("enabled").MustBool(false),
		},
	}

	if hs.Cfg.GeomapDefaultBaseLayerConfig != nil {
		jsonObj["geomapDefaultBaseLayerConfig"] = hs.Cfg.GeomapDefaultBaseLayerConfig
	}
	if !hs.Cfg.GeomapEnableCustomBaseLayers {
		jsonObj["geomapDisableCustomBaseLayer"] = true
	}

	return jsonObj, nil
}

func getPanelSort(id string) int {
	sort := 100
	switch id {
	case "timeseries":
		sort = 1
	case "barchart":
		sort = 2
	case "stat":
		sort = 3
	case "gauge":
		sort = 4
	case "bargauge":
		sort = 5
	case "table":
		sort = 6
	case "singlestat":
		sort = 7
	case "piechart":
		sort = 8
	case "state-timeline":
		sort = 9
	case "heatmap":
		sort = 10
	case "status-history":
		sort = 11
	case "histogram":
		sort = 12
	case "graph":
		sort = 13
	case "text":
		sort = 14
	case "alertlist":
		sort = 15
	case "dashlist":
		sort = 16
	case "news":
		sort = 17
	}
	return sort
}

func (hs *HTTPServer) GetFrontendSettings(c *models.ReqContext) {
	settings, err := hs.getFrontendSettingsMap(c)
	if err != nil {
		c.JsonApiErr(400, "Failed to get frontend settings", err)
		return
	}

	c.JSON(200, settings)
}

// EnabledPlugins represents a mapping from plugin types (panel, data source, etc.) to plugin IDs to plugins
// For example ["panel"] -> ["piechart"] -> {pie chart plugin DTO}
type EnabledPlugins map[plugins.Type]map[string]plugins.PluginDTO

func (ep EnabledPlugins) Get(pluginType plugins.Type, pluginID string) (plugins.PluginDTO, bool) {
	if _, exists := ep[pluginType][pluginID]; exists {
		return ep[pluginType][pluginID], true
	}

	return plugins.PluginDTO{}, false
}

func (hs *HTTPServer) enabledPlugins(ctx context.Context, orgID int64) (EnabledPlugins, error) {
	ep := make(EnabledPlugins)

	pluginSettingMap, err := hs.pluginSettings(ctx, orgID)
	if err != nil {
		return ep, err
	}

	apps := make(map[string]plugins.PluginDTO)
	for _, app := range hs.pluginStore.Plugins(ctx, plugins.App) {
		if b, exists := pluginSettingMap[app.ID]; exists {
			app.Pinned = b.Pinned
			apps[app.ID] = app
		}
	}
	ep[plugins.App] = apps

	dataSources := make(map[string]plugins.PluginDTO)
	for _, ds := range hs.pluginStore.Plugins(ctx, plugins.DataSource) {
		if _, exists := pluginSettingMap[ds.ID]; exists {
			dataSources[ds.ID] = ds
		}
	}
	ep[plugins.DataSource] = dataSources

	panels := make(map[string]plugins.PluginDTO)
	for _, p := range hs.pluginStore.Plugins(ctx, plugins.Panel) {
		if _, exists := pluginSettingMap[p.ID]; exists {
			panels[p.ID] = p
		}
	}
	ep[plugins.Panel] = panels

	return ep, nil
}

func (hs *HTTPServer) pluginSettings(ctx context.Context, orgID int64) (map[string]*models.PluginSettingInfoDTO, error) {
	pluginSettings := make(map[string]*models.PluginSettingInfoDTO)

	// fill settings from database
	if pss, err := hs.SQLStore.GetPluginSettings(ctx, orgID); err != nil {
		return nil, err
	} else {
		for _, ps := range pss {
			pluginSettings[ps.PluginId] = ps
		}
	}

	// fill settings from app plugins
	for _, plugin := range hs.pluginStore.Plugins(ctx, plugins.App) {
		// ignore settings that already exist
		if _, exists := pluginSettings[plugin.ID]; exists {
			continue
		}

		// add new setting which is enabled depending on if AutoEnabled: true
		pluginSetting := &models.PluginSettingInfoDTO{
			PluginId: plugin.ID,
			OrgId:    orgID,
			Enabled:  plugin.AutoEnabled,
			Pinned:   plugin.AutoEnabled,
		}

		pluginSettings[plugin.ID] = pluginSetting
	}

	// fill settings from all remaining plugins (including potential app child plugins)
	for _, plugin := range hs.pluginStore.Plugins(ctx) {
		// ignore settings that already exist
		if _, exists := pluginSettings[plugin.ID]; exists {
			continue
		}

		// add new setting which is enabled by default
		pluginSetting := &models.PluginSettingInfoDTO{
			PluginId: plugin.ID,
			OrgId:    orgID,
			Enabled:  true,
		}

		// if plugin is included in an app, check app settings
		if plugin.IncludedInAppID != "" {
			// app child plugins are disabled unless app is enabled
			pluginSetting.Enabled = false
			if p, exists := pluginSettings[plugin.IncludedInAppID]; exists {
				pluginSetting.Enabled = p.Enabled
			}
		}
		pluginSettings[plugin.ID] = pluginSetting
	}

	return pluginSettings, nil
}
