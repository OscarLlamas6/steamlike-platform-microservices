import { Router } from 'express'
import { addRegion, getRegion, getRegionsList, updateRegion, deleteRegion } from '../controllers/region.controller'

const router = Router();

router.post("/create", addRegion);
router.get("/:idRegion", getRegion);
router.get("/list", getRegionsList);
router.put("/update", updateRegion);
router.delete("/delete/:idRegion", deleteRegion);

export default router;