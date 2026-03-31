package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest05 {
    @Test
    public void testCpuLoadNotNaN() {
        double load = CpuHelper.getSystemCpuLoad();
        assertFalse(Double.isNaN(load));
    }

    @Test
    public void testParallelAndExecutor05() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(40_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(400);
        assertTrue(ex >= 0);
    }
}
