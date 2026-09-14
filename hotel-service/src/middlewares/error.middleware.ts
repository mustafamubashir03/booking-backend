import { NextFunction, Request, Response } from 'express';
import { AppError } from '../utils/errors/app.error';

export const appErrorHandler = (err: AppError | Error, req: Request, res: Response, next: NextFunction) => {
  console.log(err);

  const statusCode = 'statusCode' in err ? (err as AppError).statusCode : 500;
  res.status(statusCode).json({
    success: false,
    message: err.message,
  });
};

export const genericErrorHandler = (
  err: Error,
  req: Request,
  res: Response,
  next: NextFunction,
) => {
  console.log(err);

  res.status(500).json({
    success: false,
    message: 'Internal Server Error',
  });
};
