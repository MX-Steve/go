#!/usr/bin/env bash
echo "manual" > /sys/class/drm/card0/device/power_dpm_force_performance_level
echo "s 0 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "s 1 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "s 2 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "s 3 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "s 4 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "s 5 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltagea
echo "s 6 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "s 7 1100 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "m 2 2000 875" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo "c" > /sys/class/drm/card0/device/pp_od_clk_voltage
echo 3 > /sys/class/drm/card0/device/pp_dpm_sclk
echo 2 > /sys/class/drm/card0/device/pp_dpm_mclk
