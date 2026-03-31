package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest06 {
    @Test
    public void testCpuLoadBounds() {
        double load = CpuHelper.getSystemCpuLoad();
        assertTrue(load >= -1.0);
        assertTrue(load <= 1.0);
    }

    @Test
    public void testParallelAndExecutor06() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(30_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(300);
        assertTrue(ex >= 0);
    }
}
