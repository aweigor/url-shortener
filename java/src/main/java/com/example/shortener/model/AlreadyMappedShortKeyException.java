package com.example.shortener.model;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.BAD_REQUEST)
public class AlreadyMappedShortKeyException extends ModelException {
    public AlreadyMappedShortKeyException(String key, Throwable cause) {
        super(key, cause, true, false);
    }
}
