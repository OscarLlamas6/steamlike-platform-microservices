import { Router } from 'express'
import { addGameToMyLibrary, getUserLibraryGames, updateLibraryGame, deleteLibraryGame, removeGameFromLibrary } from '../controllers/library.controller'

const router = Router();

router.post("/create", addGameToMyLibrary);
router.get("/:username", getUserLibraryGames);
router.put("/update", updateLibraryGame);
router.delete("/delete/:idMyGame", deleteLibraryGame);
router.delete("/remove/:idMyGame", removeGameFromLibrary);

export default router;