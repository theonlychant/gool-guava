package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest08 {
    @Test
    public void testAvailableProcessorsReasonable() {
        int p = CpuHelper.getAvailableProcessors();
        assertTrue(p < 1024);
    }

    @Test
    public void testParallelAndExecutor08() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(25_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(250);
        assertTrue(ex >= 0);
    }
}
