package com.example.shortener.model;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.BAD_REQUEST)
public class InvalidUrlException extends ModelException {
    public InvalidUrlException(String request) {
        super(request);
    }
}
