import { Module } from '@nestjs/common';
import { ProfesoresModule } from './profesores/profesores.module';
import { CursosModule } from './cursos/cursos.module';

@Module({
  imports: [ProfesoresModule, CursosModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
