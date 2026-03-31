package com.example.guava;

import org.junit.Test;
import static org.junit.Assert.*;

import com.google.common.cache.Cache;
import com.google.common.collect.ImmutableList;

public class GuavaExamplesTest {
    @Test
    public void testImmutableList() {
        ImmutableList<String> l = GuavaExamples.makeImmutableList("a", "b", "c");
        assertEquals(3, l.size());
        assertEquals("a", l.get(0));
    }

    @Test
    public void testCache() {
        Cache<String, String> c = GuavaExamples.buildSimpleCache();
        c.put("k", "v");
        assertEquals("v", c.getIfPresent("k"));
    }

    @Test
    public void testThrottle() {
        GuavaExamples.throttleExample();
        // just ensure no exception thrown
        assertTrue(true);
    }
}
