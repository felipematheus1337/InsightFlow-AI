package spring_ai.v1.client.dto;

import java.util.List;

public record EmployeeListResponse(String message, List<EmployeeResponse> data) {
}
