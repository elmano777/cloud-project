import { Controller, Get, Post, Body, Patch, Param, Delete, ParseIntPipe } from '@nestjs/common';
import { EjerceService } from './ejerce.service';
import { CreateEjerceDto } from './dto/create-ejerce.dto';
import { UpdateEjerceDto } from './dto/update-ejerce.dto';

@Controller('ejerce')
export class EjerceController {
  constructor(private readonly ejerceService: EjerceService) {}

  @Post()
  create(@Body() createEjerceDto: CreateEjerceDto) {
    return this.ejerceService.create(createEjerceDto);
  }

  @Get()
  findAll() {
    return this.ejerceService.findAll();
  }

  @Get(':id')
  findOne(@Param('id', ParseIntPipe) id: number) {
    return this.ejerceService.findOne(id);
  }

  @Patch(':id')
  update(
    @Param('id', ParseIntPipe) id: number,
    @Body() updateEjerceDto: UpdateEjerceDto,
  ) {
    return this.ejerceService.update(id, updateEjerceDto);
  }

  @Delete(':id')
  remove(@Param('id', ParseIntPipe) id: number) {
    return this.ejerceService.remove(id);
  }
}
