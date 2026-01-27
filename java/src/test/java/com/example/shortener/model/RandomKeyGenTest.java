package com.example.shortener.model;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.core.IsNot.not;

@SpringBootTest
@AutoConfigureMockMvc
class RandomKeyGenTest {

    @Autowired
    private RandomKeyGen gen;

    @Test
    void shouldGenerateRandomKey() throws Exception {
        String keyOne = gen.generateKey(10);
        String keyTwo = gen.generateKey(10);
        assertThat(keyOne, not(keyTwo));
    }
}