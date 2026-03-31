package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest09 {
    @Test
    public void testLoadNotLessThanMinusOne() {
        assertTrue(CpuHelper.getSystemCpuLoad() >= -1.0);
    }

    @Test
    public void testParallelAndExecutor09() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(15_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(150);
        assertTrue(ex >= 0);
    }
}
