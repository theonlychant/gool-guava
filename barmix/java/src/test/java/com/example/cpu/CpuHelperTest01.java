package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest01 {
    @Test
    public void testAvailableProcessors() {
        int p = CpuHelper.getAvailableProcessors();
        assertTrue(p >= 1);
    }

    @Test
    public void testParallelAndExecutor01() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(100_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(1000);
        assertTrue(ex >= 0);
    }
}
