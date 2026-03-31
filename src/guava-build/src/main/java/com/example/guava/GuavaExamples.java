package com.example.guava;

import com.google.common.collect.ImmutableList;
import com.google.common.cache.Cache;
import com.google.common.cache.CacheBuilder;
import com.google.common.util.concurrent.RateLimiter;

import java.util.concurrent.TimeUnit;

public final class GuavaExamples {
    private GuavaExamples() {}

    public static ImmutableList<String> makeImmutableList(String... items) {
        return ImmutableList.copyOf(items);
    }

    public static Cache<String, String> buildSimpleCache() {
        return CacheBuilder.newBuilder()
                .maximumSize(100)
                .expireAfterWrite(10, TimeUnit.MINUTES)
                .build();
    }

    public static void throttleExample() {
        RateLimiter limiter = RateLimiter.create(10.0); // 10 permits per second
        // acquire a permit before doing work
        limiter.acquire();
    }
}
