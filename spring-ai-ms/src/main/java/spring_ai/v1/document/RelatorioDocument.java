package spring_ai.v1.document;

import java.time.Instant;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "relatorios")
public class RelatorioDocument {

    @Id
    private String id;
    private String relatorio;
    private Instant geradoEm;
    private String status;

    public RelatorioDocument() {
    }

    public RelatorioDocument(String relatorio, Instant geradoEm, String status) {
        this.relatorio = relatorio;
        this.geradoEm = geradoEm;
        this.status = status;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getRelatorio() {
        return relatorio;
    }

    public void setRelatorio(String relatorio) {
        this.relatorio = relatorio;
    }

    public Instant getGeradoEm() {
        return geradoEm;
    }

    public void setGeradoEm(Instant geradoEm) {
        this.geradoEm = geradoEm;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

}
