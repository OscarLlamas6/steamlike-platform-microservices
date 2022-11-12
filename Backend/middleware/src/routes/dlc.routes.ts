import { Router } from 'express'
import { 
    addDLC, 
    getDLC,
    getDLCNoAuth, 
    getDLCAdmin, 
    getDLCList, 
    getDLCListNoAuth, 
    getAllDLC, 
    updateDLC, 
    deleteDLC, 
    restoreDLC,  
    getDLCListByGame,
    getDLCListByGameNoAuth
} from '../controllers/dlc.controller'

const router = Router();

router.post("/create", addDLC);
router.get("/single/:idDLC", getDLC);
router.get("/single/noauth/:idDLC", getDLCNoAuth);
router.get("/info/:idDLC", getDLCAdmin);
router.get("/restore/:idDLC", restoreDLC);
router.get("/list", getDLCList);
router.get("/game/list/:idGame", getDLCListByGame);
router.get("/game/list/:idGame/noauth", getDLCListByGameNoAuth);
router.get("/list/noauth", getDLCListNoAuth);
router.get("/list/all", getAllDLC);
router.put("/update", updateDLC);
router.delete("/delete/:idDLC", deleteDLC);

export default router;