package question;

public class Picture {
    Piece[][] piecess;

    class Piece {
        //方块的四边形状 平：0 凸：1 凹：-1
        Integer up;
        Integer down;
        Integer left;
        Integer right;
    }

    public boolean checkPicture(Picture picture) {
        // 判空
        if (picture == null || picture.piecess == null) {
            return  false;
        }

        int rows = piecess.length;
        int columns = piecess[0].length;
        for (int i = 0; i < rows; i++) {
            if (columns != piecess[i].length) {
                return false;
            }
            for (int j = 0; j < columns; j++) {
                Piece piece = piecess[i][j];
                if (piece == null) {
                    return false;
                }

                // 最外层边缘校验
                if (i == 0 && piece.up != 0) {
                    return false;
                }

                if (i == rows - 1 && piece.down != 0) {
                    return false;
                }

                if (j == 0 && piece.left != 0) {
                    return false;
                }

                if (j == columns - 1 && piece.right != 0) {
                    return false;
                }

                // 遍历校验right和down，对比下一个右边和下一个下面的边缘
                if (j < columns - 1 && !matchRight(piece,piecess[i][j+1])) {
                    return false;
                }

                if (i < rows - 1 && !matchDown(piece,piecess[i+1][j])) {
                    return false;
                }
            }
        }
        return true;
    }

    public boolean matchRight(Piece piece1, Piece piece2) {
        if (piece1.right == -1 && piece2.left != 1) {
            return false;
        }

        if (piece1.right == 1 && piece2.left != -1) {
            return false;
        }

        return true;
    }

    public boolean matchDown(Piece piece1, Piece piece2) {
        if (piece1.down == -1 && piece2.up != 1) {
            return false;
        }

        if (piece1.down == 1 && piece2.up != -1) {
            return false;
        }

        return true;
    }
}
