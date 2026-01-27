package com.example.shortener.model;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.NOT_FOUND)
public class RedirectionNotFoundException extends ModelException {
    public RedirectionNotFoundException(String request) {
        super(request);
    }
}
