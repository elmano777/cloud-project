import { IsInt, IsNotEmpty, IsOptional, IsDateString } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class CreateEjerceDto {
  @ApiProperty({
    description: 'ID del profesor asignado',
    example: 1,
  })
  @IsInt()
  @IsNotEmpty()
  profesor_id: number;

  @ApiProperty({
    description: 'ID del curso asignado',
    example: 2,
  })
  @IsInt()
  @IsNotEmpty()
  curso_id: number;

  @ApiProperty({
    description: 'Fecha de asignación del profesor al curso',
    example: '2025-05-10T00:00:00Z',
    required: false,
  })
  @IsOptional()
  @IsDateString()
  fecha_asignacion?: string;
}
