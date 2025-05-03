import { IsInt, IsNotEmpty, IsOptional } from 'class-validator';

export class CreateEjerceDto {
  @IsInt()
  @IsNotEmpty()
  profesor_id: number;

  @IsInt()
  @IsNotEmpty()
  curso_id: number;

  @IsOptional()
  @IsInt()
  fecha_asignacion?: Date;
}
