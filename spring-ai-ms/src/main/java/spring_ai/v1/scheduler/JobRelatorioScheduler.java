package spring_ai.v1.scheduler;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import spring_ai.v1.service.RelatorioJobService;

@Component
public class JobRelatorioScheduler {

    private static final Logger log = LoggerFactory.getLogger(JobRelatorioScheduler.class);

    private final RelatorioJobService relatorioJobService;

    public JobRelatorioScheduler(RelatorioJobService relatorioJobService) {
        this.relatorioJobService = relatorioJobService;
    }

    @Scheduled(cron = "${app.scheduler.relatorio.cron}")
    public void executar() {
        log.info("Iniciando job de geração de relatório...");
        try {
            relatorioJobService.executar();
            log.info("Job de geração de relatório finalizado com sucesso.");
        } catch (Exception e) {
            log.error("Falha ao executar job de geração de relatório", e);
        }
    }

}
