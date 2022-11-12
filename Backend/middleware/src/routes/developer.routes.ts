import { Router } from 'express'
import { addDeveloper, getDeveloper, getDeveloperNoAuth, getDevelopersList, getDevelopersListNoAuth, updateDeveloper, deleteDeveloper, getDevelopersListV2 } from '../controllers/developer.controller'

const router = Router();

router.get("/", getDevelopersListV2);
router.post("/create", addDeveloper);
router.get("/:idDeveloper", getDeveloper);
router.get("/noauth/:idDeveloper", getDeveloperNoAuth);
router.get("/list", getDevelopersList);
router.get("/list/noauth", getDevelopersListNoAuth);
router.put("/update", updateDeveloper);
router.delete("/delete/:idDeveloper", deleteDeveloper);

export default router;