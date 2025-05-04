import { IsInt, IsNotEmpty, IsOptional, IsDateString } from 'class-validator';

export class CreateEjerceDto {
  @IsInt()
  @IsNotEmpty()
  profesor_id: number;

  @IsInt()
  @IsNotEmpty()
  curso_id: number;

  @IsOptional()
  @IsDateString()
  fecha_asignacion?: string;
}
