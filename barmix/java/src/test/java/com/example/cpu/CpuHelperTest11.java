package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest11 {
    @Test
    public void testMultipleAvailableProcessors() {
        int a = CpuHelper.getAvailableProcessors();
        int b = CpuHelper.getAvailableProcessors();
        assertEquals(a, b);
    }

    @Test
    public void testParallelAndExecutor11() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(11000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(110);
        assertTrue(ex >= 0);
    }
}
