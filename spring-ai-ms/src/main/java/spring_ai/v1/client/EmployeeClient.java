package spring_ai.v1.client;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;

import spring_ai.v1.client.dto.EmployeeListResponse;

@FeignClient(name = "employee-api", url = "${app.employee.api.url}")
public interface EmployeeClient {

    @GetMapping("/api/v1/employees/")
    EmployeeListResponse listarFuncionarios();

}
