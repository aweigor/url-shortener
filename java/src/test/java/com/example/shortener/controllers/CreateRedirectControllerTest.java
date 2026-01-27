package com.example.shortener.controllers;

import com.example.shortener.model.RandomKeyGen;
import com.example.shortener.services.UrlValidator;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import java.util.HashMap;
import static org.springframework.http.MediaType.APPLICATION_JSON;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
@AutoConfigureMockMvc
public class CreateRedirectControllerTest {
    @Autowired
    private MockMvc mockMvc;

    @MockBean
    UrlValidator validator;

    @MockBean
    private RandomKeyGen gen;

    static String jsonBodyParams (String longUrl) throws JsonProcessingException {
        HashMap<String, String> map = new HashMap<>();
        map.put("longUrl", longUrl);
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writer().writeValueAsString(map);
    }

    static void initGen (RandomKeyGen gen) {
        RandomKeyGen gen2 = new RandomKeyGen();
        String shortKey = gen2.generateKey(3);
        String secret = gen2.generateKey(10);
        Mockito.when(gen.generateKey(3)).thenReturn(shortKey);
        Mockito.when(gen.generateKey(10)).thenReturn(secret);
    }

    @Test
    void shouldBeOk() throws Exception {
        initGen(gen);
        String longUrl = "https://ya.ru";
        Mockito.when(validator.validateAndGetError(longUrl)).thenReturn(null);
        this.mockMvc.perform( post ("/make_shorter")
                .contentType(APPLICATION_JSON)
                .content(jsonBodyParams(longUrl)))
                .andDo(print())
                .andExpect(status().isOk());
    }
    @Test
    void shouldReturnErrorIfShortKeyAlreadyMapped() throws Exception {
        initGen(gen);
        String longUrl = "https://ya.ru";
        Mockito.when(validator.validateAndGetError(longUrl)).thenReturn(null);
        this.mockMvc.perform( post ("/make_shorter")
                    .contentType(APPLICATION_JSON)
                    .content(jsonBodyParams(longUrl)))
                .andDo(print());
        this.mockMvc.perform( post ("/make_shorter")
                .contentType(APPLICATION_JSON)
                .content(jsonBodyParams(longUrl)))
                .andDo(print())
                .andExpect(status().is4xxClientError());
    }

    @Test
    void shouldReturnErrorIfUrlIsNotValid() throws Exception {
        initGen(gen);
        String longUrl = "https://ya.ru";
        Mockito.when(validator.validateAndGetError(longUrl)).thenReturn("error");
        this.mockMvc.perform( post ("/make_shorter")
                    .contentType(APPLICATION_JSON)
                    .content(jsonBodyParams(longUrl)))
                .andDo(print())
                .andExpect(status().is4xxClientError());
    }
}
