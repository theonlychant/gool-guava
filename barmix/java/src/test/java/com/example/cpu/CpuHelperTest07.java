package com.example.cpu;

import org.junit.Test;
import static org.junit.Assert.*;

public class CpuHelperTest07 {
    @Test
    public void testMainRuns() {
        CpuHelper.main(new String[0]);
        assertTrue(CpuHelper.getAvailableProcessors() >= 0);
    }

    @Test
    public void testParallelAndExecutor07() throws Exception {
        long sum = ConcurrencyExamples.parallelSum(20_000);
        assertTrue(sum > 0);
        long ex = ConcurrencyExamples.executorWork(200);
        assertTrue(ex >= 0);
    }
}
