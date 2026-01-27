package com.example.shortener.controllers;

import com.example.shortener.messages.CreateRedirectResponse;
import com.example.shortener.model.RandomKeyGen;
import com.example.shortener.model.Redirection;
import com.example.shortener.model.RedirectionRepo;
import com.example.shortener.services.UrlShortenerService;
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
import org.springframework.test.web.servlet.MvcResult;

import java.util.HashMap;
import java.util.Optional;

import static org.hamcrest.core.StringContains.containsString;
import static org.springframework.http.MediaType.APPLICATION_JSON;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;


@SpringBootTest
@AutoConfigureMockMvc
class AdminControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private RedirectionRepo repo;

    @MockBean
    UrlValidator validator;

    static String jsonBodyParams (HashMap<String, Object> params) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writer().writeValueAsString(params);
    }

    @Test
    void shouldReturnErrorIfRedirectionNotFound() throws Exception {
        Optional<Redirection> redirection = Optional.empty();
        Mockito.when(repo.findBySecretKey(null)).thenReturn(redirection);

        this.mockMvc.perform(get("/admin/null"))
                .andDo(print())
                .andExpect(status().is4xxClientError());
    }

    @Test
    void shouldDeleteRedirection() throws Exception {
        String longUrl = "https://ya.ru";
        HashMap<String, Object> params = new HashMap<>();
        params.put("longUrl", longUrl);

        Mockito.when(validator.validateAndGetError(longUrl)).thenReturn(null);
        MvcResult result = this.mockMvc.perform( post ("/make_shorter")
                    .contentType(APPLICATION_JSON)
                    .content(jsonBodyParams(params)))
                .andDo(print())
                .andReturn();

        String json = result.getResponse().getContentAsString();
        ObjectMapper mapper = new ObjectMapper();
        CreateRedirectResponse response = mapper.readValue(json, CreateRedirectResponse.class);

        String secret = response.getSecretKey();

        Optional<Redirection> redirection = Optional.of(Mockito.mock(Redirection.class));
        Mockito.when(repo.findBySecretKey(secret)).thenReturn(redirection);

        this.mockMvc.perform(delete(String.format("/admin/%s", secret)))
                .andDo(print())
                .andExpect(status().isOk());
    }

    @Test
    void shodulReturnAdminStats() throws Exception {
        String longUrl = "https://ya.ru";
        ObjectMapper mapper = new ObjectMapper();
        String json;
        MvcResult result;

        HashMap<String, Object> params = new HashMap<>();
        params.put("longUrl", longUrl);

        Mockito.when(validator.validateAndGetError(longUrl)).thenReturn(null);
        result = this.mockMvc.perform( post ("/make_shorter")
                        .contentType(APPLICATION_JSON)
                        .content(jsonBodyParams(params)))
                .andDo(print())
                .andReturn();

        json = result.getResponse().getContentAsString();
        CreateRedirectResponse response = mapper.readValue(json, CreateRedirectResponse.class);
        String secret = response.getSecretKey();

        Optional<Redirection> redirection = Optional.of(Mockito.mock(Redirection.class));
        Mockito.when(repo.findBySecretKey(secret)).thenReturn(redirection);

        HashMap<String, Object> expectedParams = new HashMap<>();
        expectedParams.put("usageCount", 0);
        expectedParams.put("creationDate", null);

        this.mockMvc.perform(get(String.format("/admin/%s", secret)).contentType(APPLICATION_JSON))
                .andDo(print())
                .andExpect(status().isOk())
                .andExpect(content().json(jsonBodyParams(expectedParams)));
    }

}