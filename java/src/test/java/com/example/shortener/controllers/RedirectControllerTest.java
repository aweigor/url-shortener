package com.example.shortener.controllers;

import com.example.shortener.messages.CreateRedirectResponse;
import com.example.shortener.model.Redirection;
import com.example.shortener.model.RedirectionRepo;
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

import static org.springframework.http.MediaType.APPLICATION_JSON;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest
@AutoConfigureMockMvc
public class RedirectControllerTest {
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
    void shouldRedirect() throws Exception {
        String longUrl = "https://ya.ru";
        Redirection redirection = new Redirection(longUrl, null, null);
        Optional<Redirection> optional = Optional.of(redirection);

        Mockito.when(repo.findByShortKey("aaa")).thenReturn(optional);

        this.mockMvc.perform(get("/aaa"))
                .andDo(print())
                .andExpect(redirectedUrl(longUrl))
                .andExpect(status().is3xxRedirection());

    }
}
