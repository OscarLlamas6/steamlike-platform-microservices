import { Router } from 'express'
import { addGame, getGameByID, getGameByIDNoAuth, getGameAdmin, getGamesList, getGamesListV2, getGamesListByAgeRestriction, getAllGames, getGameForHomePage, updateGame, deleteGame, restoreGame } from '../controllers/game.controller'

const router = Router();

router.get("/", getGamesListV2);
router.post("/create", addGame);
router.get("/single/:idGame", getGameByID);
router.get("/single/noauth/:idGame", getGameByIDNoAuth);
router.get("/info/:idGame", getGameAdmin);
router.get("/restore/:idGame", restoreGame);
router.get("/list", getGamesList);
router.get("/list/age", getGamesListByAgeRestriction);
router.get("/list/all", getAllGames);
router.get("/list/home", getGameForHomePage);
router.put("/update", updateGame);
router.delete("/delete/:idGame", deleteGame);

export default router;