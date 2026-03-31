package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest10 {
    @Test
    public void testLoadReturnType() {
        double d = CpuHelper.getSystemCpuLoad();
        assertNotNull(d);
    }

    @Test
    public void testParallelAndExecutor10() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(12000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(120);
        assertTrue(ex >= 0);
    }
}
