package spring_ai.v1.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import spring_ai.v1.document.RelatorioDocument;

public interface RelatorioRepository extends MongoRepository<RelatorioDocument, String> {
}
