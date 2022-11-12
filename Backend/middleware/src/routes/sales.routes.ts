import { Router } from 'express'
import { addSale, getSaleByID, getSalesByUsername, getAllSales, updateSale, deleteDLC } from '../controllers/sales.controller'

const router = Router();

router.post("/create", addSale);
router.get("/single/:idSale", getSaleByID);
router.get("/list/:username", getSalesByUsername);
router.get("/list/all", getAllSales);
router.put("/update", updateSale);
router.delete("/delete/:idSale", deleteDLC);

export default router;