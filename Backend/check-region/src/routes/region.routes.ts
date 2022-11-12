import { Router } from 'express'
import { checkRegion } from '../controllers/region.controller'

const router = Router();

router.get("/", checkRegion);

export default router;