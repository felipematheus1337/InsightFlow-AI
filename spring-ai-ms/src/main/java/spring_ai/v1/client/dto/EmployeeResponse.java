package spring_ai.v1.client.dto;

import java.time.Instant;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;

public record EmployeeResponse(
        String name,
        int age,
        @JsonProperty("created_at") Instant createdAt,
        Instant updatedAt,
        Object deletedAt,
        List<String> tasks
) {
}
