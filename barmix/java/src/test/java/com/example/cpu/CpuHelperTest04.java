package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest04 {
    @Test
    public void testMultipleCalls() {
        for (int i = 0; i < 5; i++) {
            int p = CpuHelper.getAvailableProcessors();
            assertTrue(p >= 1);
        }
    }

    @Test
    public void testParallelAndExecutor04() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(60_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(600);
        assertTrue(ex >= 0);
    }
}
