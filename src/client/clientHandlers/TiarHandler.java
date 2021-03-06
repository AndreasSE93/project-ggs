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


import packageManaging.Encoder;
import packageManaging.KickMessage;
import packageManaging.Message;
import packageManaging.StageFlipper;
import packageManaging.StartableMessage;
import packageManaging.StartedMessage;
import packageManaging.TiarUserMessage;


import tiar.TiarGUI;
import clientNetworking.NetManager;

/* Handler and initializer for a Tic Tac Toe game */

public class TiarHandler implements HandlerInterface, ActionListener, MouseListener {
	NetManager network;
	TiarGUI tg;

	
	Encoder enc = new Encoder();
	StageFlipper saveMsg;
	int Player;
	final String userName;
	public boolean loop;
	private boolean startable, started;

	public TiarHandler(NetManager net, String username) {
		this.network = net;
		this.userName = username;
		this.loop = true;
		this.startable = false;
		this.started = false;

	}

	public StageFlipper init(StageFlipper sf) {
		tg = new TiarGUI();
		tg.render(this.userName);
		this.saveMsg = sf;

		for (int i = 0; i < tg.game.length; i++) {
			tg.game[i].addMouseListener(this);

		}
		tg.startGame.addActionListener(this);

		tg.chat.field.addActionListener(this);
		tg.leaveGame.addActionListener(this);

		tg.window.setVisible(true);
		return runTicTac();
	}
	
	public StageFlipper runTicTac() {
		if(saveMsg != null) {
			@SuppressWarnings("unused")
			StageFlipper startup = saveMsg;
			//Use the message when starting gameroom
			saveMsg = null;
		}
		while(loop) {
			try {
				String mess = receiveMessage();

				int id = retrieveId(mess);
				decodeAndRender(id, mess);

			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
		tg.window.setVisible(false);
		return saveMsg;
	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		
		case 100: // Chat message
			Message chatMessage = enc.chattMessageDecode(message);
			tg.chat.chatUpdate(chatMessage.message, chatMessage.user);
			break;
		
		case 201: // Move message
			TiarUserMessage mess = enc.tiarUserMessageDecode(message);
			if (this.started) {
				if (mess.isValid == 1) {
					tg.updateGameBoard(mess.Gameboard);
					tg.gl.changeTurn();
				} else
					tg.updateGameBoard(mess.Gameboard);
			
				if(mess.HasWon != 0){
					JOptionPane.showMessageDialog(null, "Player: " + Integer.toString(mess.HasWon) + " has won!", "Winner!", JOptionPane.INFORMATION_MESSAGE);
					tg.clearBoard();
				}
				if (mess.IsDraw ==1 ){
					JOptionPane.showMessageDialog(null, "The game is a draw!", "Draw!", JOptionPane.INFORMATION_MESSAGE);
					tg.clearBoard();
				}
				
			}
			break;
			
		case 202:
			StartableMessage tsm = enc.startebleMessageDecode(message);
			this.startable = tsm.isStartable;
			break;
			
		case 203:
			StartedMessage startedGame = enc.startedMessageDecode(message);
			this.started = startedGame.started;
			this.Player = startedGame.playerID;
			
			break;
			
		case 404:
			this.saveMsg = new StageFlipper(enc.kickMessageDecode(message));
			this.loop = false;
			break;

		default: 
			System.out.println(id + "\nstring: " + message);
			break;

		}

	}

	public void sendMessage(String message) {
		network.send(message);

	}

	public String receiveMessage() throws IOException {
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

				JSONtext = enc.encode(new Message(tg.chat.field.getText(),
						tg.chat.userName));
				tg.chat.field.setText("");
				sendMessage(JSONtext);
			} catch (JSONException e1) {
				e1.printStackTrace();
			}
			break;
		case "startButton":
			if(startable){
				try {
					JSONtext = enc.encode(new StartedMessage());
					sendMessage(JSONtext);
				} catch (JSONException e1) {
					e1.printStackTrace();
				}
				
				
			}
		break;
		case "kick":
			
				try {
					JSONtext = enc.encode(new KickMessage());
					sendMessage(JSONtext);
				} catch (JSONException e1) {
					e1.printStackTrace();
				}
				
				
			
		break;
		default:
	
		}
		
	}

	@Override
	public void mouseClicked(MouseEvent arg0) {

		if (Player == tg.gl.getTurn()) {
			JButton l = (JButton) arg0.getSource();
			int move = tg.getInt(l.getName());
			
			if (tg.gl.validMove(move, this.Player)) {
				TiarUserMessage tum = new TiarUserMessage(move, this.Player);
				try {
					sendMessage(enc.encode(tum));
				} catch (JSONException e) {
					e.printStackTrace();
				}
			}
		} else {
			JOptionPane.showMessageDialog(null, "Not your turn!",
					"Warning", JOptionPane.ERROR_MESSAGE);
		}
	}

	@Override
	public void mouseEntered(MouseEvent arg0) {


	}

	@Override
	public void mouseExited(MouseEvent arg0) {


	}

	@Override
	public void mousePressed(MouseEvent arg0) {


	}

	@Override
	public void mouseReleased(MouseEvent arg0) {


	}

}
