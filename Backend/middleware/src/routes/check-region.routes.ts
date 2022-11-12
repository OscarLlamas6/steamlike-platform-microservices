import { Router } from 'express'
import { checkRegion } from '../controllers/check-region.controller'

const router = Router();

router.get("/", checkRegion);

export default router;