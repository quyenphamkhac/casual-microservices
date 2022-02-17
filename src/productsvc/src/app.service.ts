import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';

@Injectable()
export class AppService {
  constructor(@Inject('PRODUCT_SERVICE') private client: ClientProxy) {}

  async publishEvent(pattern: any, data: any) {
    this.client.emit(pattern, data);
  }
}
