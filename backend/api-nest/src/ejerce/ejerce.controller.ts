import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  ParseIntPipe,
  Query,
} from '@nestjs/common';
import { EjerceService } from './ejerce.service';
import { CreateEjerceDto } from './dto/create-ejerce.dto';
import { UpdateEjerceDto } from './dto/update-ejerce.dto';
import {
  ApiTags,
  ApiOperation,
  ApiResponse,
  ApiParam,
  ApiBody,
} from '@nestjs/swagger';
import { EjerceEntity, DeleteEjerceResponse } from './dto/response.models';

@ApiTags('ejerce')
@Controller('ejerce')
export class EjerceController {
  constructor(private readonly ejerceService: EjerceService) {}

  @Post()
  @ApiOperation({ summary: 'Crear una nueva asignación de profesor a curso' })
  @ApiBody({ type: CreateEjerceDto })
  @ApiResponse({
    status: 201,
    description: 'La asignación ha sido creada exitosamente.',
    type: EjerceEntity,
  })
  @ApiResponse({ status: 400, description: 'Datos inválidos.' })
  create(@Body() createEjerceDto: CreateEjerceDto) {
    return this.ejerceService.create(createEjerceDto);
  }

  @Get()
  @ApiOperation({ summary: 'Obtener todas las asignaciones' })
  @ApiResponse({
    status: 200,
    description: 'Lista de todas las asignaciones.',
    type: EjerceEntity,
  })
  findAll(@Query('limit') limit = 10, @Query('page') page = 1) {
    return this.ejerceService.findAll(+limit, +page);
  }

  @Get(':id')
  @ApiOperation({ summary: 'Obtener una asignación por ID' })
  @ApiParam({ name: 'id', description: 'ID de la asignación' })
  @ApiResponse({
    status: 200,
    description: 'La asignación fue encontrada exitosamente.',
    type: EjerceEntity,
  })
  @ApiResponse({ status: 404, description: 'Asignación no encontrada.' })
  findOne(@Param('id', ParseIntPipe) id: number) {
    return this.ejerceService.findOne(id);
  }

  @Patch(':id')
  @ApiOperation({ summary: 'Actualizar una asignación existente' })
  @ApiParam({ name: 'id', description: 'ID de la asignación' })
  @ApiBody({ type: UpdateEjerceDto })
  @ApiResponse({
    status: 200,
    description: 'La asignación ha sido actualizada exitosamente.',
    type: EjerceEntity,
  })
  @ApiResponse({ status: 404, description: 'Asignación no encontrada.' })
  @ApiResponse({ status: 400, description: 'Datos inválidos.' })
  update(
    @Param('id', ParseIntPipe) id: number,
    @Body() updateEjerceDto: UpdateEjerceDto,
  ) {
    return this.ejerceService.update(id, updateEjerceDto);
  }

  @Delete(':id')
  @ApiOperation({ summary: 'Eliminar una asignación' })
  @ApiParam({ name: 'id', description: 'ID de la asignación' })
  @ApiResponse({
    status: 200,
    description: 'La asignación ha sido eliminada exitosamente.',
    type: DeleteEjerceResponse,
  })
  @ApiResponse({ status: 404, description: 'Asignación no encontrada.' })
  remove(@Param('id', ParseIntPipe) id: number) {
    return this.ejerceService.remove(id);
  }
}
