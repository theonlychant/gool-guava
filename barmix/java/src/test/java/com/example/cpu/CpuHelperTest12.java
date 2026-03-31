package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest12 {
    @Test
    public void testLoadWithinBounds() {
        double load = CpuHelper.getSystemCpuLoad();
        assertTrue(load <= 1.0);
        assertTrue(load >= -1.0);
    }

    @Test
    public void testParallelAndExecutor12() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(10000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(100);
        assertTrue(ex >= 0);
    }
}
