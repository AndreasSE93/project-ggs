package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;
import java.io.IOException;

import javax.swing.JButton;
import javax.swing.JOptionPane;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.ChatMessageEncoder;
import packageManaging.Message;
import packageManaging.TiarStartMessage;
import packageManaging.TiarStartMessageEncoder;
import packageManaging.TiarUserMessage;
import packageManaging.TiarUserMessageEncoder;

import tiar.TiarGUI;

import clientNetworking.NetManager;

/* Handler and initializer for a Tic Tac Toe game */

public class TiarHandler implements HandlerInterface, ActionListener,
		MouseListener {

	NetManager network;
	TiarGUI tg;
	ChatMessageEncoder cme = new ChatMessageEncoder();
	TiarUserMessageEncoder tume = new TiarUserMessageEncoder();
	TiarStartMessageEncoder tsme = new TiarStartMessageEncoder();
	int Player = 2;
	final String userName;
	private boolean loop = true;

	public TiarHandler(NetManager net, String username) {
		this.network = net;
		this.userName = username;

	}

	public int init() {
		tg = new TiarGUI();
		tg.render(this.userName);

		for (int i = 0; i < tg.game.length; i++) {
			tg.game[i].addMouseListener(this);

		}

		tg.chat.field.addActionListener(this);

		tg.window.setVisible(true);

		while (loop) {
			try {
				String mess = recieveMessage();

				int id = retrieveId(mess);
				decodeAndRender(id, mess);

			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
		return 1;
	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		case 100: // Chat message
			System.out.println(message);
			Message chatMessage = cme.decode(message);
			tg.chat.chatUpdate(chatMessage.message, chatMessage.user);
			break;

		case 201: // Move message

			TiarUserMessage mess = tume.decode(message);
			if (mess.isValid == 1) {
				tg.updateGameBoard(mess.Gameboard);
				tg.gl.changeTurn();
			} else
				tg.updateGameBoard(mess.Gameboard);
			
			if(mess.HasWon != 0){
				JOptionPane.showMessageDialog(null, "Player: " + Integer.toString(mess.HasWon) + " has won!", "Winner!", JOptionPane.ERROR_MESSAGE);
				tg.clearBoard();
			}
			if (mess.IsDraw ==1 ){
				tg.clearBoard();
			}
			
			break;

		case 202:

			TiarStartMessage tsm = tsme.decode(message);
			Player = tsm.turn;

			break;

		default: // Should not come here
				System.out.println( id + "\nstring: " + message);
			break;

		}

	}

	public void sendMessage(String message) {
		network.send(message);

	}

	public String recieveMessage() throws IOException {
		return network.receiveMessage();
	}

	public int retrieveId(String mess) {
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {

		}
		return 0;

	}

	public void actionPerformed(ActionEvent e) {
		String command = e.getActionCommand();
		String JSONtext = "";
		switch (command) {
		case "chatmessage": // JTextField
			try {

				JSONtext = cme.encode(new Message(tg.chat.field.getText(),
						tg.chat.userName));
				tg.chat.field.setText("");
			} catch (JSONException e1) {
				e1.printStackTrace();
			}
			break;
		default:

			break;
		}
		sendMessage(JSONtext);
	}

	@Override
	public void mouseClicked(MouseEvent arg0) {
		if (Player == tg.gl.getTurn()) {
			JButton l = (JButton) arg0.getSource();

			int move = tg.getInt(l.getName());
			if (tg.gl.validMove(move, this.Player)) {
				TiarUserMessage tum = new TiarUserMessage(move, this.Player);
				try {
					sendMessage(tume.encode(tum));
				} catch (JSONException e) {
					// TODO Auto-generated catch block
					e.printStackTrace();
				}
			}
		} else {
			JOptionPane.showMessageDialog(null, "Not your turn!", "Warning",
					JOptionPane.ERROR_MESSAGE);
		}

	}

	public void checkWinner() {

		switch (tg.gl.hasWon()) {
		case 1:
			System.out.println("Spelare 1 vann");
			/* Bör skriva ut ett meddelande */
			tg.clearBoard();
			break;
		case 2:
			System.out.println("Spelare 2 vann");
			/* Bör skriva ut ett meddelande */
			tg.clearBoard();
			break;
		default:
			if (tg.gl.isDraw() == 1) {
				tg.clearBoard();
			}
			// System.out.println(tg.gl.toString());
		}

	}

	@Override
	public void mouseEntered(MouseEvent arg0) {
		// TODO Auto-generated method stub

	}

	@Override
	public void mouseExited(MouseEvent arg0) {
		// TODO Auto-generated method stub

	}

	@Override
	public void mousePressed(MouseEvent arg0) {
		// TODO Auto-generated method stub

	}

	@Override
	public void mouseReleased(MouseEvent arg0) {
		// TODO Auto-generated method stub

	}

}
