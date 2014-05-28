package snake;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.GridLayout;
import java.awt.Image;
import java.awt.image.BufferedImage;

import javax.swing.ImageIcon;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextArea;
import javax.swing.border.EmptyBorder;

public class SnakeGUI {

	public JFrame window = new JFrame();
	ImageIcon ii;
	BufferedImage bi;
	public JLabel gamePane;
	Image dimg;
	public JPanel achtungPanel;
	public JButton leaveGame;
	public JButton startGame;

	final private int WIDTH = 1035 - 200;
	final private int HEIGTH = 790;

	private JTextArea p1, p2, p3, p4, p5, p6, p7, p8;
	public JTextArea[] playerScores = { p1, p2, p3, p4, p5, p6, p7, p8 };

	private int[] ColorArray = { Color.red.getRGB(), Color.blue.getRGB(),
			Color.green.getRGB(), Color.yellow.getRGB(), Color.pink.getRGB(), Color.cyan.getRGB(), Color.lightGray.getRGB(), Color.orange.getRGB(), Color.white.getRGB() };

	public int nameSet = 0;

	public void render() {
		window.setSize(1035, 790);
	

		window.setLocationRelativeTo(null);
		window.setResizable(false);
		window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		window.setBackground(Color.DARK_GRAY);

		achtungPanel = new JPanel(new BorderLayout());
		achtungPanel.setPreferredSize(new Dimension(1035, 790));

		gamePane = new JLabel();
		gamePane.setFocusable(true);


		achtungPanel.add(createTextPanel(), BorderLayout.EAST);
		achtungPanel.add(gamePane);

		bi = new BufferedImage(WIDTH, HEIGTH, BufferedImage.TYPE_INT_RGB);
		gamePane.setFocusable(true);
		renderNewGame();

		window.add(achtungPanel);

		window.setVisible(true);
		window.validate();



	}

	public JPanel createTextPanel() {

		JPanel gameInfo = new JPanel(new BorderLayout());
		gameInfo.setPreferredSize(new Dimension(200, 790));
		gameInfo.setBackground(Color.gray.darker());

		JTextArea gameName = new JTextArea("          Achtung Die Kurve");
		Font font = new Font("Helvetica", Font.BOLD, 12);
		gameName.setFont(font);
		gameName.setEditable(false);
		gameName.setBackground(Color.gray.darker());
		gameName.setForeground(Color.WHITE);
		gameName.setBorder(new EmptyBorder(0, 0, 200, 0));
		gameInfo.add(gameName, BorderLayout.NORTH);

		createPlayerPanels(gameInfo);
		JPanel buttons = new JPanel();

		leaveGame = new JButton("Leave");
		leaveGame.setFont(font);
		leaveGame.setBackground(Color.gray.darker());
		leaveGame.setActionCommand("kick");

		startGame = new JButton("Start");
		startGame.setFont(font);
		startGame.setBackground(Color.gray.darker());
		startGame.setActionCommand("startButton");

		buttons.setBackground(Color.gray.darker());
		buttons.add(startGame);
		buttons.add(leaveGame);
		gameInfo.add(buttons, BorderLayout.SOUTH);

		return gameInfo;

	}

	public void createPlayerPanels(JPanel gameInfo) {
		Font playerFont = new Font("Helvetica", Font.BOLD, 20);
		JPanel panelScores = new JPanel(new GridLayout(8,1));
		panelScores.setBorder(new EmptyBorder(0, 40, 0, 0));
		panelScores.setBackground(Color.gray.darker());
		int i;
		for (i = 0; i < playerScores.length; i++) {

			playerScores[i] = new JTextArea("Player " + i + ": 0");
			playerScores[i].setVisible(false);
			playerScores[i].setFont(playerFont);
			playerScores[i].setBackground(Color.gray.darker());
			

			
			Color clr = new Color(ColorArray[i]);
			playerScores[i].setForeground(clr);
			panelScores.add(playerScores[i]);

		}
		gameInfo.add(panelScores,BorderLayout.CENTER);

	}

	public void updateScore(SnakePlayer p) {
		int id = p.PlayerID - 1;
		if (id >= 0) {
			String stringText = playerScores[id].getText();
		
		int index = 0;
		for (int i = stringText.length() - 1; i >= 0; i--) {
			if (stringText.substring(i - 1, i).equals(":")) {
				index = i;
				break;
			}
		
		}

		playerScores[id].setText(stringText.substring(0, index) + " "
				+ p.getScore());
		}
	}

	public void repaint(SnakePlayer[] players) {
		for (int i = 0; i < players.length; i++) {
			if (players[i].isAlive()) {

				updateScore(players[i]);
				for (int d = 0; d < 10; d = d + 2) {
					 int player = players[i].PlayerID - 1;
					 if (player >= 0){
						 
					 
					bi.setRGB((int) players[i].playerArray[d],
							(int) players[i].playerArray[d + 1],
							ColorArray[players[i].PlayerID - 1]);
					ii = new ImageIcon();
					ii.setImage(bi);

					gamePane.setIcon(ii);
					gamePane.validate();
				}
				ii = new ImageIcon();
				ii.setImage(bi);

				gamePane.setIcon(ii);
				gamePane.validate();
				 }
			}
		}
	}

	public void setNames(SnakeServerMessage SSM) {
		for (int i = 0; i < SSM.Players.length; i++) {
			String name = SSM.Players[i].PlayerName;
			if (name.length() < 2) {

				name = "Nils   " + i;
			}
			if (name.length() < 5) {

				name += "   ";
			}
			if (name.length() > 8) {
				name = name.substring(0, 7);
			}
			playerScores[i].setText(name + ": " + SSM.Players[i].getScore());
		}
		showScores(SSM);
		nameSet = 1;
	}

	public void renderNewGame() {



		for (int k = 0; k < WIDTH; k++) {
			for (int i = 0; i < HEIGTH; i++) {
				bi.setRGB(k, i, Color.BLACK.getRGB());

			}
		}
		ii = new ImageIcon();
		ii.setImage(bi);

		gamePane.setIcon(ii);
		gamePane.validate();

	}



	public void showScores(SnakeServerMessage SSM) {
		
		
		for (int i = 0 ; i < SSM.Players.length; i++){
			playerScores[i].setVisible(true);
			playerScores[i].setEditable(false);
		}
	}

}
