import { Module } from '@nestjs/common';
import { UsuariosModule } from './usuarios/usuarios.module';
import { EjerceModule } from './ejerce/ejerce.module';

@Module({
  imports: [UsuariosModule, EjerceModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
