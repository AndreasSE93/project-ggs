package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;
import java.io.IOException;

import javax.swing.JButton;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.ChatMessageEncoder;
import packageManaging.Message;
import packageManaging.TiarUserMessage;
import packageManaging.TiarUserMessageEncoder;


import tiar.TiarGUI;

import clientNetworking.NetManager;



/* Handler and initializer for a Tic Tac Toe game */

public class TiarHandler implements HandlerInterface, ActionListener, MouseListener{

	NetManager network;
	TiarGUI tg;
	ChatMessageEncoder cme = new ChatMessageEncoder();
	TiarUserMessageEncoder tume = new TiarUserMessageEncoder();
	final String userName;
	private boolean loop = true;
	
	public TiarHandler(NetManager net , String username){
		this.network = net;
		this.userName = username;
		
		
	}

	public int init(){
		 tg = new TiarGUI();
		 tg.render(this.userName);
		
		 System.out.println("1");
		 for (int i = 0 ; i < tg.game.length; i++){
		 tg.game[i].addMouseListener(this);
		
		 
		 }
		 System.out.println("2");
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
		switch(id){
		case 100: //Chat message
			ChatMessageEncoder mh = new ChatMessageEncoder();
			Message chatMessage = mh.decode(message);
			tg.chat.chatUpdate(chatMessage.message, chatMessage.user);
				break;
				
		case 200: //Move message
				TiarUserMessageEncoder s = new TiarUserMessageEncoder();
				TiarUserMessage mess = s.decode(message);
				tg.doMove(mess.Move , tg.gl.getTurn());
				switch (tg.gl.hasWon()){
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
					if(tg.gl.isDraw() == 1){tg.clearBoard();}
					System.out.println(tg.gl.toString());
				}
				break;
				
		default: //Should not come here
				break;
		
		
		}
		
		
		
	}
	
	public void sendMessage(String message) {
		network.send(message);
		
	}

	
	public String recieveMessage() throws IOException {
		return network.receiveMessage();
	}
	
	
	public int retrieveId(String mess){
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
		case "chatmessage":  //JTextField
			try {
				
				JSONtext = cme.encode(new Message(tg.chat.field.getText(), tg.chat.userName));
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
		JButton l =(JButton) arg0.getSource();
		String name = l.getName();
		tg.doMove(name, tg.gl.getTurn());
		System.out.println(name);
		TiarUserMessage tum = new TiarUserMessage(tg.gl.isDraw(), tg.gl.hasWon(), name);
		try {
			sendMessage(tume.encode(tum));
		} catch (JSONException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		switch (tg.gl.hasWon()){
		case 1:
			
			break;
		case 2:
			
			
			break;
		default:
			
			
			break;
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
