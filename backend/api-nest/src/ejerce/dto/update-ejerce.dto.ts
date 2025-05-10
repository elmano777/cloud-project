import { IsInt, IsOptional, IsDateString } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class UpdateEjerceDto {
  @ApiProperty({
    description: 'ID del profesor asignado',
    example: 1,
    required: false,
  })
  @IsOptional()
  @IsInt()
  profesor_id?: number;

  @ApiProperty({
    description: 'ID del curso asignado',
    example: 2,
    required: false,
  })
  @IsOptional()
  @IsInt()
  curso_id?: number;

  @ApiProperty({
    description: 'Fecha de asignación del profesor al curso',
    example: '2025-05-10T00:00:00Z',
    required: false,
  })
  @IsOptional()
  @IsDateString()
  fecha_asignacion?: string;
}
