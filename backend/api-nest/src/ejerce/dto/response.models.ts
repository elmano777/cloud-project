import { ApiProperty } from '@nestjs/swagger';

export class EjerceEntity {
  @ApiProperty({ example: 1 })
  id: number;

  @ApiProperty({ example: 1 })
  profesor_id: number;

  @ApiProperty({ example: 2 })
  curso_id: number;

  @ApiProperty({ example: '2025-05-10T12:00:00Z' })
  fecha_asignacion: Date;
}

export class DeleteEjerceResponse {
  @ApiProperty({ example: 'Registro ejerce eliminado correctamente' })
  message: string;

  @ApiProperty({ example: 1 })
  ejerceEliminado: number;
}
