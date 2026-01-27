package com.example.shortener.controllers;

import com.example.shortener.model.Redirection;
import com.example.shortener.services.UrlShortenerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import jakarta.servlet.http.HttpServletResponse;

@Controller
public class RedirectController {

    @Autowired
    UrlShortenerService shortener;

    @RequestMapping("/{shortKey}")
    public void doRedirect(@PathVariable String shortKey, HttpServletResponse response) {
        Redirection redirection = shortener.resolve(shortKey);
        response.setStatus(HttpServletResponse.SC_MOVED_TEMPORARILY);
        response.setHeader("Location", redirection.getLongUrl());
    }

}
