package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest02 {
    @Test
    public void testSystemCpuLoadRange() {
        double load = CpuHelper.getSystemCpuLoad();
        assertTrue(load >= -1.0 && load <= 1.0);
    }

    @Test
    public void testParallelAndExecutor02() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(50_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(800);
        assertTrue(ex >= 0);
    }
}
