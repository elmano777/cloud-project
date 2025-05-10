import { ApiProperty } from '@nestjs/swagger';

export class UsuarioEntity {
  @ApiProperty({ example: 1 })
  id: number;

  @ApiProperty({ example: 'Juan' })
  nombre: string;

  @ApiProperty({ example: 'Pérez' })
  apellido: string;

  @ApiProperty({ example: 'juan.perez@ejemplo.com' })
  email: string;

  @ApiProperty({ example: 'profesor' })
  rol: string;

  @ApiProperty({ example: '2025-05-10T12:00:00Z' })
  fechaCreacion: Date;

  @ApiProperty({ example: '2025-05-10T12:00:00Z' })
  fechaActualizacion: Date;
}

export class DeleteUsuarioResponse {
  @ApiProperty({ example: 'Usuario eliminado correctamente' })
  message: string;

  @ApiProperty({ example: 'juan.perez@ejemplo.com' })
  usuarioEliminado: string;
}
