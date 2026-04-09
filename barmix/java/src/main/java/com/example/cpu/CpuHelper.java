package com.example.cpu;

import java.lang.management.ManagementFactory;

/**
 * Minimal CpuHelper used by barmix tests.
 */
public final class CpuHelper {
    private CpuHelper() {}

    public static int getAvailableProcessors() {
        return Runtime.getRuntime().availableProcessors();
    }

    public static double getSystemCpuLoad() {
        try {
            java.lang.management.OperatingSystemMXBean os = ManagementFactory.getOperatingSystemMXBean();
            // try to use com.sun.management.OperatingSystemMXBean if available
            if (os instanceof com.sun.management.OperatingSystemMXBean) {
                com.sun.management.OperatingSystemMXBean sun = (com.sun.management.OperatingSystemMXBean) os;
                double v = sun.getSystemCpuLoad();
                return v;
            }
        } catch (Throwable t) {
            // ignore and fallthrough
        }
        return -1.0;
    }

    public static void main(String[] args) {
        System.out.println("Processors: " + getAvailableProcessors());
        System.out.println("System CPU load: " + getSystemCpuLoad());
    }
}
