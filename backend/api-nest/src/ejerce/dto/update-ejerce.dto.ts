import { IsInt, IsOptional, IsDateString } from 'class-validator';

export class UpdateEjerceDto {
  @IsOptional()
  @IsInt()
  profesor_id?: number;

  @IsOptional()
  @IsInt()
  curso_id?: number;

  @IsOptional()
  @IsDateString()
  fecha_asignacion?: Date;
}
