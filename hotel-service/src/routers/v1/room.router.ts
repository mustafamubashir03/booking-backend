import express from "express"

import { roomGenerationController } from "../../controllers/roomGeneration.controller"
import { validateRequestBody } from "../../validators";
import { roomGenerationJobSchema } from "../../validators/roomGeneration.validator";
import { createRoomCategorySchema } from "../../validators/roomCategory.validator";
import { roomCategoryController } from "../../controllers/roomCategory.controller";


const roomRouter = express.Router();

roomRouter.post('/generate', validateRequestBody(roomGenerationJobSchema), roomGenerationController)
roomRouter.post('/room-category', validateRequestBody(createRoomCategorySchema), roomCategoryController)

export default roomRouter