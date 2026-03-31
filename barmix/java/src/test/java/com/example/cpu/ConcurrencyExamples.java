package com.example.cpu;

import java.util.concurrent.*;
import java.util.stream.IntStream;

public final class ConcurrencyExamples {
    private ConcurrencyExamples() {}

    // Parallel stream example: sum of 1..n using parallel streams
    public static long parallelSum(int n) {
        return IntStream.rangeClosed(1, n)
                .parallel()
                .mapToLong(i -> i)
                .sum();
    }

    // Executor example: split work into tasks and run with a fixed thread pool
    public static long executorWork(int totalItems) throws InterruptedException, ExecutionException {
        int threads = Runtime.getRuntime().availableProcessors();
        ExecutorService ex = Executors.newFixedThreadPool(threads);
        try {
            int chunk = Math.max(1, totalItems / threads);
            Callable<Long> worker = () -> {
                long local = 0;
                for (int i = 0; i < chunk; i++) {
                    // small CPU-bound work
                    int v = i + 1;
                    local += (long) v * v;
                }
                return local;
            };

            Future<Long>[] futures = new Future[threads];
            for (int t = 0; t < threads; t++) {
                futures[t] = ex.submit(worker);
            }

            long total = 0;
            for (Future<Long> f : futures) {
                total += f.get();
            }
            return total;
        } finally {
            ex.shutdownNow();
        }
    }
}
