// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     TSTypesJenny
//     PluginTsTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

import * as ui from '@grafana/schema';

export interface Options extends ui.OptionsWithLegend, ui.OptionsWithTooltip, ui.OptionsWithTimezones {
  /**
   * Controls the column width
   */
  colWidth?: number;
  /**
   * Enables pagination when > 0
   */
  perPage?: number;
  /**
   * Set the height of the rows
   */
  rowHeight: number;
  /**
   * Show values on the columns
   */
  showValue: ui.VisibilityMode;
}

export const defaultOptions: Partial<Options> = {
  colWidth: 0.9,
  perPage: 20,
  rowHeight: 0.9,
  showValue: ui.VisibilityMode.Auto,
};

export interface FieldConfig extends ui.AxisConfig, ui.HideableFieldConfig {
  fillOpacity?: number;
  lineWidth?: number;
}

export const defaultFieldConfig: Partial<FieldConfig> = {
  fillOpacity: 70,
  lineWidth: 1,
};
