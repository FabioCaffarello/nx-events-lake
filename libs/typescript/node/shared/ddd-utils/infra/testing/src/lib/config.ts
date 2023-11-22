import { config as readEnv } from 'dotenv';
import { join } from 'path';

export class Config {
  static env: any = null;

  static db(service = 'admin-videos-catalog') {
    Config.readEnv(service);

    return {
      dialect: 'sqlite' as any,
      host: Config.env.DB_HOST,
      logging: Config.env.DB_LOGGING === 'true',
    };
  }

  static readEnv(service = 'admin-videos-catalog') {
    if (Config.env) {
      return;
    }

    console.log('NODE_ENV', process.env['NODE_ENV']);
    console.log("__dirname", __dirname);
    Config.env = readEnv({
      path: join(__dirname, `../../../../../../../../../../apps/client-layer/backend/${service}/envs/.env.${process.env["NODE_ENV"]}`),
    }).parsed;
  }
}
