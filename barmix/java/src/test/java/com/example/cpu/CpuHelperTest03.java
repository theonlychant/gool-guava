package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest03 {
    @Test
    public void testAvailableProcessorsPositive() {
        assertTrue(CpuHelper.getAvailableProcessors() > 0);
    }

    @Test
    public void testParallelAndExecutor03() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(75_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(900);
        assertTrue(ex >= 0);
    }
}
